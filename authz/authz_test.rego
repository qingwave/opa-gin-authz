package authz

test_get_non_api {
    allow with input as {"path": "/", "method": "GET"}
}

test_post_non_api {
    not allow with input as {"path": "/", "method": "POST"}
}

test_authenticated_get_api {
    allow with input as {"path": "/api/some", "method": "GET", "subject": {"user": "someone"}}
}

test_unauthenticated_get_api {
    not allow with input as {"path": "/api/some", "method": "GET"}
}
