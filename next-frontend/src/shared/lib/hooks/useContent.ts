"use client";

import { ReactNode, useState } from "react";

export const useContent = (defaultContent: ReactNode) => {
  const [currentContent, setCurrentContent] =
    useState<ReactNode>(defaultContent);

  const setContent = (newContent: ReactNode) => {
    setCurrentContent(newContent);
  };

  const resetContent = () => {
    setCurrentContent(defaultContent);
  };

  return [currentContent, setContent, resetContent] as const;
};
