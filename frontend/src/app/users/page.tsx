"use client";

import { apiRequest } from "@/utils/api";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type User = {
	id: string;
	name: string;
};

export default function UsersPage() {
	const [users, setUsers] = useState<User[]>([]);
	const [error, setError] = useState<string | null>(null);
	const router = useRouter();

	useEffect(() => {
		const fetchUsers = async () => {
			try {
				const response = await apiRequest<User[]>({
					method: "GET",
					url: "/users",
				});
				setUsers(response);
			} catch (err) {
				setError("Failed to fetch users.");
				// You might want to handle authorization differently here
				// For now, just redirecting to login
				router.push("/login");
			}
		};

		fetchUsers();
	}, [router]);

	if (error) {
		return <div>{error}</div>;
	}

	return (
		<div>
			<h1>User List</h1>
			<ul>
				{users.map((user) => (
					<li key={user.id}>
						{user.name} (ID: {user.id})
					</li>
				))}
			</ul>
		</div>
	);
}
