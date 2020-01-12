# Introduction to Ansible
Hello there, I'm following tutorial from acloud.guru that demonstrate the basic features on Ansible. In this tutorial, I learned how to setup 2 tiers of PHP web server & a load balancer (using Apache2) on Debian distro.

The source code for this lesson is available here: `https://github.com/ACloudGuru-Resources/Course_Introduction_to_Ansible`

For Ansible Roles, see the folder `../ansible-roles` in this repository

## What I have learned:
- Set up Ansible config
- Set up Ansible Inventory
- Run adhoc ansible commands/module
- Write ansible playbook
- Write ansible template file (for config file generation)
- Ansible Variables
- Ansible Galaxy, and the repository of Ansible Roles, SUPER COOL!
- Error handling - use `ignore_errors` flag to except the errors
- Ansible Vault - creates encrypted variable
- Ansible Prompts - prompts user for ceratin input using `vars_prompt`

## Ansible Commands
- `ansible -m setup app1` - See the Ansible facts
- `ansible-playbook playbooks/setup-app.yaml --check` - Dry run the playbook to detect changes to be made
- `ansible-playbook playbooks/setup-app.yaml --tags upload` - Run only `upload` tag tasks
- `ansible-playbook playbooks/setup-app.yaml --skip-tags upload` - Run all tasks *EXCEPT* `upload` tag tasks
- `ansible-playbook playbooks/setup-app.yaml --ask-vault-pass` - Run playbook that decrypt Ansible Vault variables
- `ansible-vault create vars/secret-password.yaml` - Create an encrypted variable
- `ansible-galaxy roles/webservers init ` - Create ansible roles

## Personal Note
You can see the load balancer management UI at `http://<load balancer ip>/balancer-manager`
