// "use client";

// import { useState, useEffect } from "react";
// import { Switch } from "./switch";
// import { Moon, Sun } from "lucide-react";
// import { cn } from "@/shared/lib/utils";

// export function ThemeSwitch() {
//   const [theme, setTheme] = useState<"light" | "dark">("light");
//   const [mounted, setMounted] = useState(false);

//   useEffect(() => {
//     setMounted(true);
//     const savedTheme = localStorage.getItem("theme") as "light" | "dark" | null;
//     const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    
//     if (savedTheme) {
//       setTheme(savedTheme);
//     } else if (systemPrefersDark) {
//       setTheme("dark");
//     }
//   }, []);

//   useEffect(() => {
//     if (!mounted) return;
    
//     localStorage.setItem("theme", theme);
//     const html = document.documentElement;
    
//     if (theme === "dark") {
//       html.classList.add("dark");
//       html.classList.remove("light");
//     } else {
//       html.classList.add("light");
//       html.classList.remove("dark");
//     }
//   }, [theme, mounted]);

//   const toggleTheme = () => {
//     setTheme(prev => prev === "light" ? "dark" : "light");
//   };

//   if (!mounted) {
//     return (
//       <div className="flex items-center gap-2">
//         <Sun className="h-4 w-4" />
//         <Switch disabled />
//         <Moon className="h-4 w-4" />
//       </div>
//     );
//   }

//   return (
//     <div className="flex items-center gap-2">
//       <Sun className={cn(
//         "h-4 w-4 transition-colors",
//         theme === "light" ? "text-amber-500" : "text-gray-400"
//       )} />
      
//       <Switch
//         checked={theme === "dark"}
//         onCheckedChange={toggleTheme}
//         aria-label="Переключить тему"
//         className="data-[state=checked]:bg-gray-800"
//       />
      
//       <Moon className={cn(
//         "h-4 w-4 transition-colors",
//         theme === "dark" ? "text-indigo-400" : "text-gray-400"
//       )} />
//     </div>
//   );
// }

"use client";

import { useState, useEffect } from "react";
import { Switch } from "./switch";
import { Moon, Sun } from "lucide-react";
import { cn } from "@/shared/lib/utils";

export function ThemeSwitch() {
  const [theme, setTheme] = useState<"light" | "dark">("light");
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
    const savedTheme = localStorage.getItem("theme") as "light" | "dark" | null;
    const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    
    const initialTheme = savedTheme || (systemPrefersDark ? "dark" : "light");
    setTheme(initialTheme);
    document.documentElement.classList.toggle("dark", initialTheme === "dark");
  }, []);

  const toggleTheme = () => {
    const newTheme = theme === "light" ? "dark" : "light";
    setTheme(newTheme);
    localStorage.setItem("theme", newTheme);
    document.documentElement.classList.toggle("dark", newTheme === "dark");
  };

  if (!mounted) {
    return (
      <div className="relative inline-flex h-6 w-12 items-center rounded-full bg-gray-300">
        <div className="translate-x-1 h-4 w-4 rounded-full bg-white shadow-lg" />
      </div>
    );
  }

  return (
    <div className="relative inline-flex cursor-pointer" onClick={toggleTheme}>
      <div className={cn(
        "relative h-7 w-14 rounded-full transition-all duration-300",
        theme === "dark" ? "bg-gray-800" : "bg-gray-200"
      )}>
        {/* Солнце (левая иконка) */}
        <Sun className={cn(
          "absolute left-1.5 top-1/2 h-4 w-4 -translate-y-1/2 transition-all duration-300",
          theme === "light" ? "text-amber-600 opacity-100" : "text-gray-400 opacity-40"
        )} />
        
        {/* Луна (правая иконка) */}
        <Moon className={cn(
          "absolute right-1.5 top-1/2 h-4 w-4 -translate-y-1/2 transition-all duration-300",
          theme === "dark" ? "text-indigo-300 opacity-100" : "text-gray-500 opacity-40"
        )} />
        
        {/* Бегунок с иконкой */}
        <div className={cn(
          "absolute top-1/2 h-6 w-6 -translate-y-1/2 rounded-full bg-white shadow-lg transition-all duration-300",
          theme === "dark" ? "left-7" : "left-1"
        )}>
          {theme === "dark" ? (
            <Moon className="h-3.5 w-3.5 text-orange-500 absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2" />
          ) : (
            <Sun className="h-3.5 w-3.5 text-orange-500 absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2" />
          )}
        </div>
      </div>
    </div>
  );
}