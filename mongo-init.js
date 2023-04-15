db.createUser(
        {
            user: "congestion",
            pwd: "congestion",
            roles: [
                {
                    role: "readWrite",
                    db: "volvo"
                }
            ]
        }
);
