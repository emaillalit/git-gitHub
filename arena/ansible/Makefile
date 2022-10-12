SHELL=/bin/bash

ap=ansible-playbook -b -v

# Detect tag_Role_* group_var file.
ifneq ($(wildcard ./inventory/group_vars/tag_Role_$(ROLE_UNDERSCORES)),)
ROLE_FILE="inventory/group_vars/tag_Role_$(ROLE_UNDERSCORES)"
else
ROLE_FILE=""
endif

# Checks [tag_Role_$(ROLE)] or [tag_Role_$(ROLE)(.yml)] files on existence and presence
c14n-validate:
	@test -f $(ROLE_FILE) || (echo "Cannot find group_var file for role [$(ROLE)]. Should be either [group_vars/tag_Role_$(ROLE)(.yml)] or [group_vars/tag_Role_$(ROLE)]." && exit 1)

# c14n inventory is for the deploy of services with tag_Role_template_* defined
C14N_INVENTORY_HOME:=$(shell mktemp -d -t $(ROLE_UNDERSCORES).XXXXXX)
c14n-inventory:: c14n-validate
	@# Copy group_vars
	@mkdir -p $(C14N_INVENTORY_HOME)/group_vars/all
	@cp -r inventory/group_vars/tag_Role_$(ROLE_UNDERSCORES)* $(C14N_INVENTORY_HOME)/group_vars
	@cp -r inventory/group_vars/all/global_vars_* $(C14N_INVENTORY_HOME)/group_vars/all
	@# Copy dynamic EC2 inventory with quick ROLE filter
	@cp inventory/aws_ec2.yml $(C14N_INVENTORY_HOME)
	@sed -i "s/role_template/$(ROLE)/g" $(C14N_INVENTORY_HOME)/aws_ec2.yml
	@sed -i "s/region_template/$(REGION)/g" $(C14N_INVENTORY_HOME)/aws_ec2.yml

# utility inventory is for the deploy of utility services
UTILITY_INVENTORY_HOME:=$(shell mktemp -d -t $(ROLE_UNDERSCORES).XXXXXX)
utility-inventory::
	@# Copy group_vars
	@mkdir -p $(UTILITY_INVENTORY_HOME)/group_vars/all
	@cp -r inventory/group_vars/tag_Role_$(ROLE_UNDERSCORES)* $(UTILITY_INVENTORY_HOME)/group_vars
	@cp -r inventory/group_vars/all/global_vars_* $(UTILITY_INVENTORY_HOME)/group_vars/all
	@# Copy dynamic EC2 inventory with quick ROLE filter
	@cp inventory/aws_ec2.yml $(UTILITY_INVENTORY_HOME)
	@sed -i "s/role_template/$(ROLE)/g" $(UTILITY_INVENTORY_HOME)/aws_ec2.yml
	@sed -i "s/region_template/$(REGION)/g" $(UTILITY_INVENTORY_HOME)/aws_ec2.yml

# inventory for service without tag role file
NO_TAG_ROLE_INVENTORY_HOME:=$(shell mktemp -d -t $(ROLE_UNDERSCORES).XXXXXX)
no-tag-role-inventory::
	@# Copy dynamic EC2 inventory with quick ROLE filter
	@mkdir -p $(NO_TAG_ROLE_INVENTORY_HOME)/group_vars/all
	@cp -r inventory/group_vars/all/global_vars_* $(NO_TAG_ROLE_INVENTORY_HOME)/group_vars/all
	@cp inventory/aws_ec2.yml $(NO_TAG_ROLE_INVENTORY_HOME)
	@sed -i "s/role_template/$(ROLE)/g" $(NO_TAG_ROLE_INVENTORY_HOME)/aws_ec2.yml
	@sed -i "s/region_template/$(REGION)/g" $(NO_TAG_ROLE_INVENTORY_HOME)/aws_ec2.yml

c14n-deploy:: c14n-inventory
	@echo
	@echo '  ___   ___  ___  _     ___ __   __ ___  _  _   ___  '
	@echo ' |   \ | __|| _ \| |   / _ \\ \ / /|_ _|| \| | / __| '
	@echo ' | |) || _| |  _/| |__| (_) |\ V /  | | | .` || (_ | '
	@echo ' |___/ |___||_|  |____|\___/  |_|  |___||_|\_| \___| '
	@echo
	@echo 'ROLE: $(ROLE)'
	@${ap} -u ${user} -i $(C14N_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) -l $(LIMIT) deploy-c14n-pod.yml

deploy-backstop:: c14n-inventory
	@echo
	@echo "  ___   _   ___ _  _____ _____ ___  ___  "
	@echo " | _ ) /_\ / __| |/ / __|_   _/ _ \| _ \ "
	@echo " | _ \/ _ \ (__| ' <\__ \ | || (_) |  _/ "
	@echo " |___/_/ \_\___|_|\_\___/ |_| \___/|_|   "
	@echo
	@echo 'ROLE: $(ROLE)'
	@${ap} -u ${user} -i $(C14N_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) -l $(LIMIT) deploy-backstop.yml

deploy-metrics:: c14n-inventory
	@echo
	@echo '  __  __ ___ _____ ___ ___ ___ ___  '
	@echo ' |  \/  | __|_   _| _ \_ _/ __/ __| '
	@echo ' | |\/| | _|  | | |   /| | (__\__ \ '
	@echo ' |_|  |_|___| |_| |_|_\___\___|___/ '
	@echo
	@echo 'ROLE: $(ROLE)'
	@${ap} -u ${user} -i $(C14N_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) deploy-metrics.yml

deploy-fluentd:: c14n-inventory
	@echo
	@echo '    ___ _                            _  '
	@echo '   / __) |                    _     | | '
	@echo ' _| |__| | _   _ _____ ____ _| |_ __| | '
	@echo '(_   __) || | | | ___ |  _ (_   _) _  | '
	@echo '  | |  | || |_| | ____| | | || |( (_| | '
	@echo '  |_|   \_)____/|_____)_| |_| \__)____| '
	@echo
	@echo 'ROLE: $(ROLE)'
	@${ap} -u ${user} -i $(C14N_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) deploy-fluentd.yml

deploy-rabbitmq:: utility-inventory
	@echo
	@echo '  ___    _   ___ ___ ___ _____ __  __  ___   '
	@echo ' | _ \  /_\ | _ ) _ )_ _|_   _|  \/  |/ _ \  '
	@echo ' |   / / _ \| _ \ _ \| |  | | | |\/| | (_) | '
	@echo ' |_|_\/_/ \_\___/___/___| |_| |_|  |_|\__\_\ '
	@echo
	@echo 'ROLE: $(ROLE)'
	@${ap} -u ${user} -i $(UTILITY_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) deploy-rabbitmq.yml

deploy-redis:: utility-inventory
	@echo
	@echo '  ___ ___ ___ ___ ___  '
	@echo ' | _ \ __|   \_ _/ __| '
	@echo ' |   / _|| |) | |\__ \ '
	@echo ' |_|_\___|___/___|___/ '
	@echo
	@echo 'ROLE: $(ROLE)'
	@${ap} -u ${user} -i $(UTILITY_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) deploy-redis.yml

bastion-management:: c14n-inventory
	@echo
	@echo '  ___   _   ___ _____ ___ ___  _  _  '
	@echo ' | _ ) /_\ / __|_   _|_ _/ _ \| \| | '
	@echo ' | _ \/ _ \\__ \ | |  | | (_) | .` | '
	@echo ' |___/_/ \_\___/ |_| |___\___/|_|\_| '
	@echo
	@echo 'ROLE: $(ROLE)'
	@echo	
	@${ap} -u ${user} -i $(C14N_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) bastion-management.yml

deploy-mail:: c14n-inventory
	@echo
	@echo '  __  __   _   ___ _     '
	@echo ' |  \/  | /_\ |_ _| |    '
	@echo ' | |\/| |/ _ \ | || |__  '
	@echo ' |_|  |_/_/ \_\___|____| '
	@echo
	@echo 'ROLE: $(ROLE)'
	@echo
	@${ap} -u ${user} -i $(C14N_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) deploy-mail.yml

firstboot-management:: utility-inventory
	@echo
	@echo '  ___  ___  ___  ____ _____ ___   ___    ___  _____  '
	@echo ' | __||_ _|| _ \/ __||_   _| _ ) / _ \  / _ \|_   _| '
	@echo ' | _|  | | |   /\__ \  | | | _ \| (_) || (_) | | |   '
	@echo ' |_|  |___||_|_\|___/  |_| |___/ \___/  \___/  |_|   '
	@echo
	@echo 'ROLE: $(ROLE)'
	@echo
	@${ap} -u ${user} -i $(UTILITY_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) -l $(LIMIT) firstboot-management.yml

route53-dns:: utility-inventory
	@echo
	@echo '    ___  ____  __  ________________ ____    ___  _  ______ '
	@echo '   / _ \/ __ \/ / / /_  __/ __/ __/|_  /___/ _ \/ |/ / __/ '
	@echo '  / , _/ /_/ / /_/ / / / / _//__ \_/_ <___/ // /    /\ \   '
	@echo ' /_/|_|\____/\____/ /_/ /___/____/____/  /____/_/|_/___/   '
	@echo
	@echo 'ROLE: $(ROLE)'
	@echo
	@${ap} -u ${user} -i $(UTILITY_INVENTORY_HOME) -e role=$(ROLE_UNDERSCORES) route53-dns.yml

