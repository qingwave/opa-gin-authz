package authz

default allow = false

allow {
    input.method == "GET"
	not startswith(input.path, "/api")
}

allow {
    input.method == "GET"
    input.subject.user != ""
}
