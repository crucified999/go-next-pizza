"use client";

import { Provider } from "react-redux";
import { store } from "./store";
import { useCheckAuth } from "@/shared/lib/hooks/useCheckAuth";

// const AuthChecker = ({ children }: { children: React.ReactNode }) => {
//   useCheckAuth();
//   return <>{children}</>;
// };

export const ReduxProvider = ({ children }: { children: React.ReactNode }) => {
  return <Provider store={store}>{children}</Provider>;
};
