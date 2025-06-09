"use client";

import { apiRequest } from "@/utils/api";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type AuthHistory = {
	id: string;
	user_id: string;
	method: string;
	authenticated_at: string;
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
		<div className="w-full max-w-6xl mx-auto p-4">
			<h1 className="text-2xl font-bold text-blue-700 dark:text-blue-300 mb-6">認証履歴</h1>
			<div className="overflow-x-auto">
				<table className="w-full border-collapse bg-white dark:bg-gray-800 shadow-lg rounded-lg">
					<thead>
						<tr className="bg-gray-100 dark:bg-gray-700">
							<th className="px-6 py-3 text-left text-sm font-semibold text-gray-600 dark:text-gray-200">user_id</th>
							<th className="px-6 py-3 text-left text-sm font-semibold text-gray-600 dark:text-gray-200">method</th>
							<th className="px-6 py-3 text-left text-sm font-semibold text-gray-600 dark:text-gray-200">authenticated_at</th>
						</tr>
					</thead>
					<tbody>
						{history.map((entry) => (
							<tr key={entry.id} className="border-t border-gray-200 dark:border-gray-700">
								<td className="px-6 py-4 text-sm text-gray-800 dark:text-gray-200">
									{entry.user_id}
								</td>
								<td className="px-6 py-4 text-sm text-gray-800 dark:text-gray-200">
									{entry.method}
								</td>
								<td className="px-6 py-4 text-sm text-gray-800 dark:text-gray-200">
									{new Date(entry.authenticated_at).toLocaleString()}
								</td>
							</tr>
						))}
					</tbody>
				</table>
			</div>
		</div>
	);
}
