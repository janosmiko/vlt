package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"gopkg.in/yaml.v2"
)

type Config struct {
	VaultAddr         string `yaml:"vault_addr"`
	VaultNamespace    string `yaml:"vault_namespace"`
	VaultToken        string `yaml:"vault_token"`
	VaultCaCert       string `yaml:"vault_cacert"`
	VaultClientCert   string `yaml:"vault_client_cert"`
	VaultClientKey    string `yaml:"vault_client_key"`
	VltLogFile     string `yaml:"vlt_log_file"`
	VltLogLevel    string `yaml:"vlt_log_level"`
	VltRefreshRate int    `yaml:"vlt_refresh_rate"`
}

func LoadConfig(cfgFile string) Config {
	var config Config
	// Load the config from the YAML file
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory")
	}

	var data []byte
	if cfgFile == "" {
		yamlFilePath := filepath.Join(home, ".vlt.yaml")
		if _, err := os.Stat(yamlFilePath); os.IsNotExist(err) {
			fmt.Printf("Config file does not exist: %s\n", yamlFilePath)
		} else {
			data, err = os.ReadFile(yamlFilePath)
			if err != nil {
				fmt.Printf("Error reading YAML file: %v\n", err)
			}
		}
	} else {
		fmt.Println("Using config file: ", cfgFile)
		data, err = os.ReadFile(cfgFile)
		if err != nil {
			fmt.Printf("Error reading YAML file: %v\n", err)
		}
	}

	if data != nil {
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			fmt.Printf("Error parsing YAML file: %v\n", err)
		}
	}

	// Check for vault cache
	home, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory")
	} else {
		vaultTokenPath := filepath.Join(home, ".vault-token")
		if _, err := os.Stat(vaultTokenPath); os.IsNotExist(err) {
			fmt.Printf("Vault token file does not exist: %s\n", vaultTokenPath)
		} else {
			data, err := os.ReadFile(vaultTokenPath)
			if err != nil {
				fmt.Printf("Error reading vault token file: %v\n", err)
			} else {
				config.VaultToken = string(data)
			}
		}
	}

	// Overwrite with environment variables if they are set
	if vaultAddr := os.Getenv("VAULT_ADDR"); vaultAddr != "" {
		config.VaultAddr = vaultAddr
	}
	if vaultNamespace := os.Getenv("VAULT_NAMESPACE"); vaultNamespace != "" {
		config.VaultNamespace = vaultNamespace
	}
	if vaultToken := os.Getenv("VAULT_TOKEN"); vaultToken != "" {
		config.VaultToken = vaultToken
	}
	if vaultCaCert := os.Getenv("VAULT_CACERT"); vaultCaCert != "" {
		config.VaultCaCert = vaultCaCert
	}
	if vaultClientCert := os.Getenv("VAULT_CLIENT_CERT"); vaultClientCert != "" {
		config.VaultClientCert = vaultClientCert
	}
	if vaultClientKey := os.Getenv("VAULT_CLIENT_KEY"); vaultClientKey != "" {
		config.VaultClientKey = vaultClientKey
	}
	if vltLogFile := os.Getenv("VLT_LOG_FILE"); vltLogFile != "" {
		config.VltLogFile = vltLogFile
	}
	if vltLogLevel := os.Getenv("VLT_LOG_LEVEL"); vltLogLevel != "" {
		config.VltLogLevel = vltLogLevel
	}
	if vltRefreshRate := os.Getenv("VLT_REFRESH_RATE"); vltRefreshRate != "" {
		vltRefreshRateInt, err := strconv.Atoi(vltRefreshRate)
		if err != nil {
			fmt.Printf("Error converting VLT_REFRESH_RATE to int: %v", err)
		} else {
			config.VltRefreshRate = vltRefreshRateInt
		}
	}

	if config.VaultToken == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user home directory")
		} else {
			vaultTokenPath := filepath.Join(home, ".vault-token")
			if _, err := os.Stat(vaultTokenPath); err == nil {
				data, err := os.ReadFile(vaultTokenPath)
				if err != nil {
					fmt.Printf("Error reading vault token file: %v\n", err)
				} else {
					config.VaultToken = string(data)
				}
			}
		}
	}

	if config.VaultAddr == "" {
		fmt.Println("VAULT_ADDR is not set. Please set it and try again.")
		os.Exit(1)
	}

	if config.VaultToken == "" {
		fmt.Println("VAULT_TOKEN is not set. Please set it and try again.")
		os.Exit(1)
	}

	if config.VltRefreshRate == 0 {
		config.VltRefreshRate = 30
	}

	if strings.EqualFold(config.VltLogLevel, "debug") {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGTERM)

			<-ch
			fmt.Println("Dumping goroutines")
			bufsize := int(10 * 1024 * 1024) // 10 MiB
			buf := make([]byte, bufsize)
			n := runtime.Stack(buf, true)
			filename := fmt.Sprintf("%s.dump", config.VltLogFile)

			ioutil.WriteFile(filename, buf[:n], 0644)
			os.Exit(1)
		}()

	}

	return config
}
