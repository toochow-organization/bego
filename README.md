# My Backend Go workspace
This backend Go workspace will be used as a template or references from any of my future project

# Reference
This works were inspired and referenced from 
[workspace](https://github.com/anthonycorbacho/workspace) project

# Setup for Tilt
- [Rancher-desktop](https://docs.rancherdesktop.io/getting-started/installation/)
- [Tilt](https://tilt.dev/)

## Start the local dev environment 
- Run the following command in local `tilt up` and press `space` to jump to the tilt dashboard


## PostgreSQL 
- Can be access by pgadmin4, the user can be either `postgres` or `bego`, password: `password`


## Troubleshooting
- If you got some trouble with publishing image 
`Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?` especially when running in Rancher-desktop, then execute `sudo ln -s ~$USER/.rd/docker.sock /var/run/docker.sock` to map the docker.sock to /var/run/