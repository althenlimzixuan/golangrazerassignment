# Razer Assignment - Lim Zi Xuan
## Overview
This project will initiate 5 Assignment Web Service instance, 1 NGINX docker instance, and 1 MySQL docker instance in Docker. 

## Setup Guide
1. Clone the repository
2. Install Docker support in your machine. For Windows 10 Home users, remember to Enable the WSL 2 feature on Windows. (Reference: https://docs.docker.com/docker-for-windows/install/)
3. Open command prompt and navigate to cloned directory
4. In command prompt, enter following command:

   ***docker-compose up --build --remove-orphans --detach***

5. Done! Open browser and navigate to http://localhost/assets/

## Remarks
1. Currently there're only two emails registered as compromised email
    - althenlim601@gmail.com
    - zixuan_lim@yahoo.com
2. While other email will declare as safe, Email Regex validation is applied in the web page

## Risk Accessment
1. Uncertified by proper CA
   - Mitigation Plan: Need to get a domain name, generate CSR and sign off by authorized CA
2. Unencrypted configuration, may need encryption policy
   - Mitigation Plan: May need encryption policy. Eg: Sensitive credential shall be encrypted at configuration file level and decrypt in source code
