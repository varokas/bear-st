provider "google" {
  credentials = file("gcloud-account.json")
  project     = "varokas"
  region      = "us-west1"
}

resource "google_compute_project_metadata" "default" {
  metadata = {
    ssh-keys = "ubuntu:${file("gcp.pub")}"
  }
}

 resource "google_compute_instance" "service" {
   name         = "service"
   machine_type = "f1-micro"
   zone        = "us-west1-a"

   boot_disk {
     initialize_params {
       image = "ubuntu-1804-docker"
     }
   }

   network_interface {
     # A default network is created for all GCP projects
     network       = "default"
     access_config {
     }
   }
 }
