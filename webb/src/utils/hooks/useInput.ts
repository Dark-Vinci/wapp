import React, { useState } from "react";

export function useInput<T>(initialValue: T) {
  const [input, setInput] = useState<T>(initialValue);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.target.value as T);
  };

  const reset = () => {
    setInput(initialValue);
  };

  return {
    input,
    reset,
    onChange: handleChange,
  };
}
