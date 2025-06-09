"use client";
import { apiRequest } from "@/utils/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Signup() {
	const [name, setName] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState("");
	const [success, setSuccess] = useState("");
	const router = useRouter();

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		setLoading(true);
		setError("");
		setSuccess("");
		try {
			await apiRequest({
				method: "POST",
				url: "/api/signup",
				data: { name, email, password },
			});
			setSuccess("登録が完了しました。ログインページに移動します。");
			setTimeout(() => {
				router.push("/login");
			}, 2000);
		} catch (err) {
			setError("登録に失敗しました。");
		} finally {
			setLoading(false);
		}
	};

	return (
		<div className="flex flex-col w-full min-h-screen items-center justify-center bg-gradient-to-br from-white to-blue-50 dark:from-gray-900 dark:to-blue-950">
			<div className="w-1/2">
				<h1 className="text-3xl font-bold text-blue-700 dark:text-blue-300 mb-4">新規登録</h1>
				<form
					onSubmit={handleSubmit}
					className="flex flex-col gap-4 w-full bg-white dark:bg-gray-800 p-6 rounded-lg shadow"
				>
					<input
						type="text"
						placeholder="名前"
						value={name}
						onChange={(e) => setName(e.target.value)}
						required
						className="px-4 py-2 rounded border border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-100"
					/>
					<input
						type="email"
						placeholder="メールアドレス"
						value={email}
						onChange={(e) => setEmail(e.target.value)}
						required
						className="px-4 py-2 rounded border border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-100"
					/>
					<input
						type="password"
						placeholder="パスワード"
						value={password}
						onChange={(e) => setPassword(e.target.value)}
						required
						className="px-4 py-2 rounded border border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-100"
					/>
					<button
						type="submit"
						disabled={loading}
						className="px-6 py-3 rounded-lg bg-blue-600 text-white font-semibold shadow hover:bg-blue-700 transition-colors disabled:opacity-50"
					>
						{loading ? "登録中..." : "登録"}
					</button>
					{error && <p className="text-red-500 text-sm">{error}</p>}
					{success && <p className="text-green-600 text-sm">{success}</p>}
				</form>
				<p className="text-sm text-gray-600 dark:text-gray-300">
					すでにアカウントをお持ちですか？{" "}
					<Link href="/login" className="text-blue-600 dark:text-blue-300 underline">
						ログイン
					</Link>
				</p>
			</div>
		</div>
	);
}
