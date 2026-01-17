---
name: create_react_feature
description: Scaffolds a new React frontend feature following the Package by Feature pattern and Hook-based Dependency Injection.
---

# Create React Feature

This skill guides you through creating a new feature module in the React frontend.
**Target Directory**: `src/frontend/src/features/<feature_name>/`

## 1. Prerequisites
- Confirm the feature name (e.g., `accounts`, `transactions`).
- Ensure `src/frontend/src/features/` exists.

## 2. File Structure Generation
You must create the following structure:

```
src/features/<feature_name>/
├── components/      # UI Components (Presentational)
├── hooks/           # Business Logic (Containers/Services)
├── types/           # TypeScript Interfaces
└── index.ts         # Public API (Exports)
```

### A. `types/index.ts` (The Domain)
Define the data structures.
```typescript
export interface <FeatureName> {
    id: string;
    // other fields
}

export interface Create<FeatureName>Input {
    // fields
}
```

### B. `hooks/use<FeatureName>.ts` (The Service Layer)
This is where **Dependency Injection** happens. Encapsulate API calls and state management here. 
Components MUST use this hook, not fetch directly.
```typescript
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { api } from '@/lib/api'; // or axios instance

export const use<FeatureName>s = (userId: string) => {
    return useQuery({
        queryKey: ['<feature_name>', userId],
        queryFn: async () => {
            const { data } = await api.get(`/api/<feature_name>`);
            return data;
        }
    });
};

export const useCreate<FeatureName> = () => {
    const queryClient = useQueryClient();
    return useMutation({
        mutationFn: async (input: Create<FeatureName>Input) => {
            return api.post(`/api/<feature_name>`, input);
        },
        onSuccess: () => {
             queryClient.invalidateQueries({ queryKey: ['<feature_name>'] });
        }
    });
};
```

### C. `components/<FeatureName>List.tsx` (The UI)
Presentational component. Accepts data via props or uses the hook? 
*Strictly speaking*, presentational components should take props, but for rapid feature development, "Smart Components" consuming the hook are acceptable at the top level.

**Pattern:** Smart Container -> Dumb List
```typescript
import { use<FeatureName>s } from '../hooks/use<FeatureName>s';

export const <FeatureName>List = ({ userId }: { userId: string }) => {
    const { data, isLoading } = use<FeatureName>s(userId);

    if (isLoading) return <div>Loading...</div>;

    return (
        <ul>
            {data?.map(item => (
                <li key={item.id}>{item.name}</li>
            ))}
        </ul>
    );
};
```

### D. `index.ts` (The Public API)
Export only what is needed by the rest of the app.
```typescript
export * from './components/<FeatureName>List';
export * from './hooks/use<FeatureName>';
export * from './types';
```

## 3. Wiring (Manual Step)
Remind the user (or yourself) to add the new routes to `src/app/routes.tsx` or `App.tsx`:
```typescript
<Route path="/<feature_name>" element={<<FeatureName>List userId={user.id} />} />
```
