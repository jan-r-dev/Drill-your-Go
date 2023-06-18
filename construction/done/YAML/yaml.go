package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Nodes        []*Nodes      `yaml:"nodes"`
	SystemImages *SystemImages `yaml:"system_images"`
}

func (cfg *Config) modifyConfig() {
	cfg.SystemImages.Etcd = "rancher/coreos-etcd:v3.1.14"

Loop:
	for _, node := range cfg.Nodes {
		for _, role := range node.Role {
			if role == "controlplane" {
				node.Port = 2234
				break Loop
			}
		}
	}

	for _, node := range cfg.Nodes {
		if node.Address == "example.com" {
			node.InternalAddress = "192.168.1.72"
			break
		}
	}
}

type Labels struct {
	App string `yaml:"app"`
}
type Taints struct {
	Key    string `yaml:"key"`
	Value  string `yaml:"value"`
	Effect string `yaml:"effect"`
}
type Nodes struct {
	Address          string    `yaml:"address"`
	User             string    `yaml:"user"`
	Role             []string  `yaml:"role"`
	Port             int       `yaml:"port,omitempty"`
	DockerSocket     string    `yaml:"docker_socket,omitempty"`
	SSHKeyPath       string    `yaml:"ssh_key_path,omitempty"`
	SSHKey           string    `yaml:"ssh_key,omitempty"`
	SSHCertPath      string    `yaml:"ssh_cert_path,omitempty"`
	SSHCert          string    `yaml:"ssh_cert,omitempty"`
	HostnameOverride string    `yaml:"hostname_override,omitempty"`
	InternalAddress  string    `yaml:"internal_address,omitempty"`
	Labels           *Labels   `yaml:"labels,omitempty"`
	Taints           []*Taints `yaml:"taints,omitempty"`
}
type SystemImages struct {
	Kubernetes                string `yaml:"kubernetes"`
	Etcd                      string `yaml:"etcd"`
	Alpine                    string `yaml:"alpine"`
	NginxProxy                string `yaml:"nginx_proxy"`
	CertDownloader            string `yaml:"cert_downloader"`
	KubernetesServicesSidecar string `yaml:"kubernetes_services_sidecar"`
	Kubedns                   string `yaml:"kubedns"`
	Dnsmasq                   string `yaml:"dnsmasq"`
	KubednsSidecar            string `yaml:"kubedns_sidecar"`
	KubednsAutoscaler         string `yaml:"kubedns_autoscaler"`
	PodInfraContainer         string `yaml:"pod_infra_container"`
}

func main() {
	filename := "config.yaml"

	backupFile(filename)

	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}
	yaml.Unmarshal(bs, cfg)
	if err != nil {
		log.Fatal(err)
	}

	cfg.modifyConfig()

	saveYaml(cfg, filename)
}

func backupFile(filename string) {
	filepath, err := createBackupFolders()
	if err != nil {
		log.Fatal(err)
	}

	src, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(fmt.Sprintf("%s/%s", filepath, filename))
	if err != nil {
		log.Fatal(err)
	}

	bytesWritten, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Config backed up. Written %d bytes.\n", bytesWritten)
}

func createBackupFolders() (string, error) {
	ts := strconv.FormatInt(time.Now().UnixNano(), 10)
	filepath := fmt.Sprintf("backups/%v", ts)

	err := os.MkdirAll(filepath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func saveYaml(in interface{}, filename string) error {
	out, err := yaml.Marshal(in)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, out, 0644)
	if err != nil {
		return err
	} else {
		fmt.Printf("File %q updated successfully.\n", filename)
		return nil
	}
}
