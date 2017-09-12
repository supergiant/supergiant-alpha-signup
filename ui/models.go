package main

import "time"

type Session struct {
	ID        string    `json:"id"`
	UserID    *int64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`

	User string `json:"user"`
}

type HelmRelease struct {
	ChartName    string `json:"chart_name"`
	ChartVersion string `json:"chart_version"`
	Config       struct {
		API struct {
			Enabled bool `json:"enabled"`
			Image   struct {
				PullPolicy string `json:"pullPolicy"`
				Repository string `json:"repository"`
				Tag        string `json:"tag"`
			} `json:"image"`
			Name      string   `json:"name"`
			Resources struct{} `json:"resources"`
			Service   struct {
				ExternalPort int `json:"externalPort"`
				InternalPort int `json:"internalPort"`
			} `json:"service"`
			Support struct {
				Enabled  bool   `json:"enabled"`
				Password string `json:"password"`
			} `json:"support"`
		} `json:"api"`
		Ingress struct {
			Annotations struct {
				Traefik_frontend_rule_type string `json:"traefik.frontend.rule.type"`
			} `json:"annotations"`
			Enabled bool   `json:"enabled"`
			Name    string `json:"name"`
		} `json:"ingress"`
		Persistence struct {
			AccessMode   string `json:"accessMode"`
			Enabled      bool   `json:"enabled"`
			Size         string `json:"size"`
			StorageClass string `json:"storageClass"`
		} `json:"persistence"`
		UI struct {
			Enabled bool `json:"enabled"`
			Image   struct {
				PullPolicy string `json:"pullPolicy"`
				Repository string `json:"repository"`
				Tag        string `json:"tag"`
			} `json:"image"`
			Name         string   `json:"name"`
			ReplicaCount int      `json:"replicaCount"`
			Resources    struct{} `json:"resources"`
			Service      struct {
				ExternalPort int `json:"externalPort"`
				InternalPort int `json:"internalPort"`
			} `json:"service"`
		} `json:"ui"`
		Uniqueurl string `json:"uniqueurl"`
	} `json:"config"`
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	Kube      struct {
		KubeProviderString string `json:"Kube_provider_string"`
		AwsConfig          struct {
			AvailabilityZone         string   `json:"availability_zone"`
			BucketName               string   `json:"bucket_name"`
			BuildElasticFilesystem   bool     `json:"build_elastic_filesystem"`
			ElasticFilesystemID      string   `json:"elastic_filesystem_id"`
			ElasticFilesystemTargets []string `json:"elastic_filesystem_targets"`
			ElbSecurityGroupID       string   `json:"elb_security_group_id"`
			InternetGatewayID        string   `json:"internet_gateway_id"`
			LastSelectedAz           string   `json:"last_selected_az"`
			MasterVolumeSize         int      `json:"master_volume_size"`
			MultiAz                  bool     `json:"multi_az"`
			NodeSecurityGroupID      string   `json:"node_security_group_id"`
			NodeVolumeSize           int      `json:"node_volume_size"`
			PrivateKey               string   `json:"private_key"`
			PrivateNetwork           bool     `json:"private_network"`
			PublicSubnetIPRange      []struct {
				IPRange  string `json:"ip_range"`
				SubnetID string `json:"subnet_id"`
				Zone     string `json:"zone"`
			} `json:"public_subnet_ip_range"`
			Region                        string   `json:"region"`
			RouteTableID                  string   `json:"route_table_id"`
			RouteTableSubnetAssociationID []string `json:"route_table_subnet_association_id"`
			VpcID                         string   `json:"vpc_id"`
			VpcIPRange                    string   `json:"vpc_ip_range"`
			VpcManaged                    bool     `json:"vpc_managed"`
		} `json:"aws_config"`
		CloudAccountName string `json:"cloud_account_name"`
		CreatedAt        string `json:"created_at"`
		CustomFiles      string `json:"custom_files"`
		EtcdDiscoveryURL string `json:"etcd_discovery_url"`
		ExtraData        struct {
			CPULimit []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"cpu_limit"`
			CPURequest []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"cpu_request"`
			CPUUsageRate []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"cpu_usage_rate"`
			KubeCPUCapacity []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"kube_cpu_capacity"`
			KubeMemoryCapacity []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"kube_memory_capacity"`
			MemoryLimit []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"memory_limit"`
			MemoryRequest []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"memory_request"`
			MemoryUsage []struct {
				Timestamp string `json:"timestamp"`
				Value     int    `json:"value"`
			} `json:"memory_usage"`
			Metrics struct{} `json:"metrics"`
		} `json:"extra_data"`
		HeapsterMetricResolution string   `json:"heapster_metric_resolution"`
		HeapsterVersion          string   `json:"heapster_version"`
		ID                       int      `json:"id"`
		KubeMasterCount          int      `json:"kube_master_count"`
		KubernetesVersion        string   `json:"kubernetes_version"`
		MasterID                 string   `json:"master_id"`
		MasterName               string   `json:"master_name"`
		MasterNodeSize           string   `json:"master_node_size"`
		MasterNodes              []string `json:"master_nodes"`
		MasterPrivateIP          string   `json:"master_private_ip"`
		MasterPublicIP           string   `json:"master_public_ip"`
		Name                     string   `json:"name"`
		NodeSizes                []string `json:"node_sizes"`
		Password                 string   `json:"password"`
		ProviderString           string   `json:"provider_string"`
		Ready                    bool     `json:"ready"`
		ServiceString            string   `json:"service_string"`
		SSHPubKey                string   `json:"ssh_pub_key"`
		UpdatedAt                string   `json:"updated_at"`
		Username                 string   `json:"username"`
		UUID                     string   `json:"uuid"`
	} `json:"kube"`
	KubeName  string `json:"kube_name"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	RepoName  string `json:"repo_name"`
	Revision  string `json:"revision"`
	Status    struct {
		Description string `json:"description"`
		MaxRetries  int    `json:"max_retries"`
		Retries     int    `json:"retries"`
	} `json:"status"`
	StatusValue  string `json:"status_value"`
	UpdatedAt    string `json:"updated_at"`
	UpdatedValue string `json:"updated_value"`
	UUID         string `json:"uuid"`
}
