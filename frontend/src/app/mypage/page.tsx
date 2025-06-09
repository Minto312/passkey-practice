"use client";

import { apiRequest } from "@/utils/api";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type User = {
	id: string;
	name: string;
};

export default function MyPage() {
	const [user, setUser] = useState<User | null>(null);
	const [error, setError] = useState<string | null>(null);
	const router = useRouter();

	useEffect(() => {
		const fetchUser = async () => {
			try {
				const response = await apiRequest<User>({
					method: "GET",
					url: "/users/me",
				});
				setUser(response);
			} catch (err) {
				setError("Failed to fetch user data.");
				// Redirect to login page if not authenticated
				router.push("/login");
			}
		};

		fetchUser();
	}, [router]);

	if (error) {
		return <div>{error}</div>;
	}

	if (!user) {
		return <div>Loading...</div>;
	}

	return (
		<div>
			<h1>My Page</h1>
			<p>ID: {user.id}</p>
			<p>Name: {user.name}</p>
			<button type="button" onClick={() => router.push("/passkeys")}>
				Manage Passkeys
			</button>
		</div>
	);
}
