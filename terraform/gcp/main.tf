// Configure the Google Cloud provider
provider "google" {
 credentials = "${file("~/.secure/gcp.json")}"
 project     = "vpsaddict-router"
 region      = "us-east1"
}

// Terraform plugin for creating random ids
resource "random_id" "instance_id" {
 byte_length = 8
}

// A single Google Cloud Engine instance
resource "google_compute_instance" "default" {
 name         = "router-${random_id.instance_id.hex}"
 //machine_type = "f1-micro"
 machine_type = "n1-standard-1"
 zone         = "us-east1-b"

 boot_disk {
   initialize_params {
     image = "centos-cloud/centos-7"
   }
 }

 metadata_startup_script = "yum update -y; yum install -y git ansible; git clone https://github.com/jmainguy/soh.re /root/soh.re; ansible-playbook /root/soh.re/ansible/playbook.yaml"

 network_interface {
   network = "default"

   access_config {
     // Include this section to give the VM an external ip address
   }
 }
 metadata = {
    ssh-keys = "jmainguy:${file("~/.ssh/id_rsa.pub")}"
 }
}

output "ip" {
 value = "${google_compute_instance.default.network_interface.0.access_config.0.nat_ip}"
}

resource "google_compute_firewall" "default" {
 name    = "soh-router-firewall"
 network = "default"

 allow {
   protocol = "tcp"
   ports    = ["80"]
 }
}
