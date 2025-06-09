"use client";

import { apiRequest } from "@/utils/api";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type AuthHistory = {
	id: string;
	timestamp: string;
	ipAddress: string;
	userAgent: string;
};

export default function AuthHistoryPage() {
	const [history, setHistory] = useState<AuthHistory[]>([]);
	const [error, setError] = useState<string | null>(null);
	const router = useRouter();

	useEffect(() => {
		const fetchHistory = async () => {
			try {
				const response = await apiRequest<AuthHistory[]>({
					method: "GET",
					url: "/auth/history",
				});
				setHistory(response);
			} catch (err) {
				setError("Failed to fetch authentication history.");
				router.push("/login");
			}
		};

		fetchHistory();
	}, [router]);

	if (error) {
		return <div>{error}</div>;
	}

	return (
		<div>
			<h1>Authentication History</h1>
			<ul>
				{history.map((entry) => (
					<li key={entry.id}>
						<p>Timestamp: {new Date(entry.timestamp).toLocaleString()}</p>
						<p>IP Address: {entry.ipAddress}</p>
						<p>User Agent: {entry.userAgent}</p>
					</li>
				))}
			</ul>
		</div>
	);
}
