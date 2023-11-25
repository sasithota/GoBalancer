module loadbalancer/balancer

go 1.21.4

// replace corebalancer.com/balancer => ./core/balancer

replace loadbalancer.com/request => ./core/request

replace loadbalancer.com/response => ./core/response

replace loadbalancer.com/worker => ./core/worker

replace loadbalancer.com/balancer => ./core/balancer

require loadbalancer.com/balancer v0.0.0-00010101000000-000000000000

require (
	loadbalancer.com/request v0.0.0-00010101000000-000000000000 // indirect
	loadbalancer.com/response v0.0.0-00010101000000-000000000000 // indirect
	loadbalancer.com/worker v0.0.0-00010101000000-000000000000 // indirect
)
