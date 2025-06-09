import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";

// APIのベースURL（環境変数や設定ファイルから取得するのが理想）
const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";

// 共通のAPIクライアント生成関数
function createApiClient(baseURL: string = BASE_URL): AxiosInstance {
  const instance = axios.create({
    baseURL,
    headers: {
      "Content-Type": "application/json",
    },
    withCredentials: true,
  });

  instance.interceptors.request.use((config) => {
    const token = localStorage.getItem("token") || "";
    if (token) config.headers["Authorization"] = `Bearer ${token}`;
    return config;
  });

  instance.interceptors.response.use(
    (response) => response,
    (error) => {
      return Promise.reject(error);
    }
  );

  return instance;
}

export async function apiRequest<T = any, D = any>(
  config: AxiosRequestConfig<D>
): Promise<T> {
  const client = createApiClient();
  try {
    const response: AxiosResponse<T> = await client.request<T, AxiosResponse<T>, D>(config);
    return response.data;
  } catch (error: any) {
    throw error;
  }
}
