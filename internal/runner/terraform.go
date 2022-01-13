package runner

import (
	"context"
	"errors"
	"fmt"
	"github.com/datadog/stratus-red-team/internal/utils"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	"log"
)

const TerraformVersion = "1.1.2"

type TerraformManager struct {
	terraformBinaryPath string
	terraformVersion    string
	terraformHandles    map[string]*tfexec.Terraform
}

func NewTerraformManager(terraformBinaryPath string) *TerraformManager {
	fmt.Println(terraformBinaryPath)
	manager := TerraformManager{
		terraformVersion:    TerraformVersion,
		terraformHandles:    map[string]*tfexec.Terraform{},
		terraformBinaryPath: terraformBinaryPath,
	}
	manager.initialize()
	return &manager
}

func (m *TerraformManager) initialize() {
	// Download the Terraform binary if it doesn't exist already
	if !utils.FileExists(m.terraformBinaryPath) {
		terraformInstaller := &releases.ExactVersion{
			Product:                  product.Terraform,
			Version:                  version.Must(version.NewVersion(TerraformVersion)),
			InstallDir:               m.terraformBinaryPath,
			SkipChecksumVerification: false,
		}
		log.Println("Installing Terraform")
		_, err := terraformInstaller.Install(context.Background())
		if err != nil {
			log.Fatalf("error installing Terraform: %s", err)
		}
	}
}

func (m *TerraformManager) TerraformApply(directory string) error {
	terraform, err := tfexec.NewTerraform(directory, m.terraformBinaryPath)

	log.Println("Initializing Terraform")
	err = terraform.Init(context.Background())
	if err != nil {
		return errors.New("unable to initalize Terraform: " + err.Error())
	}

	log.Println("Applying Terraform")
	err = terraform.Apply(context.Background())
	if err != nil {
		return errors.New("unable to apply Terraform: " + err.Error())
	}

	return nil
}

func (m *TerraformManager) TerraformDestroy(directory string) error {
	terraform, err := tfexec.NewTerraform(directory, m.terraformBinaryPath)
	if err != nil {
		return err
	}

	log.Println("Destroying Terraform")
	return terraform.Destroy(context.Background())
}
