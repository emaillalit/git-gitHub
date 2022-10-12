rabbitmqctl add_user arena Arena12345
rabbitmqctl set_user_tags arena administrator
rabbitmqctl set_permissions -p / arena '.*' '.*' '.*'
rabbitmqctl delete_user guest