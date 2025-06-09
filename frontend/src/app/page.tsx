import Link from "next/link";

export default function Home() {
  return (
    <>
      <div className="flex flex-col items-center justify-center min-h-screen w-full bg-gradient-to-b from-blue-50 via-white to-blue-100 dark:from-gray-800 dark:via-gray-900 dark:to-gray-800">
        <section className="w-full max-w-md bg-gray-100 dark:bg-gray-700 rounded-xl shadow-lg p-8 flex flex-col items-center">
          <h1 className="text-3xl sm:text-4xl font-bold text-blue-800 dark:text-blue-100 mb-2 tracking-tight drop-shadow border-b border-blue-200 dark:border-blue-700 pb-2 w-full text-center">Passkey Practice</h1>
          <p className="text-base sm:text-lg text-gray-900 dark:text-white mb-1 mt-2">パスキーで次世代の認証体験を</p>
          <p className="text-xs text-gray-600 dark:text-gray-200">パスワードレス・セキュア・かんたんログイン</p>
        </section>
        <div className="flex flex-col sm:flex-row gap-4 mt-8">
          <Link href="/signup" className="px-8 py-4 bg-blue-600 hover:bg-blue-700 text-white font-bold rounded-lg shadow-lg transition-colors duration-200 text-center">
            新規登録
          </Link>
          <Link href="/login" className="px-8 py-4 bg-white hover:bg-gray-100 dark:bg-gray-800 dark:hover:bg-gray-700 text-blue-600 dark:text-blue-300 font-bold rounded-lg shadow-lg transition-colors duration-200 border-2 border-blue-600 dark:border-blue-300 text-center">
            ログイン
          </Link>
        </div>
      </div>
    </>
  );
}
