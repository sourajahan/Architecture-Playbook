// 1. Domain Interface (The Contract)
// We decouple the Component from the API implementation.
interface User {
  id: string;
  name: string;
}

interface UserRepository {
  getUser(id: string): Promise<User>;
}

// 2. API Implementation (The Infrastructure)
class UserApiRepository implements UserRepository {
  async getUser(id: string): Promise<User> {
    const response = await fetch(`/api/users/${id}`);
    if (!response.ok) throw new Error("Network response error");
    return response.json();
  }
}

// 3. The "Principal" Level Usage (Dependency Injection / Service)
// Note: In production, we'd inject this via Context or DI container.
export const useUser = (repo: UserRepository, id: string) => {
  // Logic is clean, testable, and independent of React/Vue/Angular
  // We can swap UserApiRepository with MockRepository in tests easily.
  return repo.getUser(id);
};
