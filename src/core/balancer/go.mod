module loadbalancer.com/balancer

go 1.21.4

replace loadbalancer.com/request => ../request

replace loadbalancer.com/response => ../response

replace loadbalancer.com/worker => ../worker

require (
	loadbalancer.com/request v0.0.0-00010101000000-000000000000
	loadbalancer.com/response v0.0.0-00010101000000-000000000000
	loadbalancer.com/worker v0.0.0-00010101000000-000000000000
)
