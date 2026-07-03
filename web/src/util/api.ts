export const api = {
  post: async (endpoint: string, data: {}) => {
    const response = await fetch(`http://localhost:8080/api${endpoint}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error(`Request failed: ${response.status}`);
    }

    return response.json();
  },
};
