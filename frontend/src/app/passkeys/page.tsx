"use client";

import { apiRequest } from "@/utils/api";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type Passkey = {
	id: string;
	publicKey: string;
	createdAt: string;
};

export default function PasskeysPage() {
	const [passkeys, setPasskeys] = useState<Passkey[]>([]);
	const [error, setError] = useState<string | null>(null);
	const router = useRouter();

	useEffect(() => {
		const fetchPasskeys = async () => {
			try {
				const response = await apiRequest<Passkey[]>({
					method: "GET",
					url: "/passkeys",
				});
				setPasskeys(response);
			} catch (err) {
				setError("Failed to fetch passkeys.");
				router.push("/login");
			}
		};

		fetchPasskeys();
	}, [router]);

	const handleAddPasskey = async () => {
		// TODO: Implement passkey creation logic
		alert("Adding a new passkey...");
	};

	const handleDeletePasskey = async (id: string) => {
		// TODO: Implement passkey deletion logic
		alert(`Deleting passkey ${id}...`);
	};

	if (error) {
		return <div>{error}</div>;
	}

	return (
		<div>
			<h1>Manage Passkeys</h1>
			<button type="button" onClick={handleAddPasskey}>
				Add New Passkey
			</button>
			<ul>
				{passkeys.map((passkey) => (
					<li key={passkey.id}>
						<span>ID: {passkey.id}</span>
						<button type="button" onClick={() => handleDeletePasskey(passkey.id)}>
							Delete
						</button>
					</li>
				))}
			</ul>
		</div>
	);
}
