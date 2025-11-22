// import { useEffect, useRef } from "react";
// import { useAppDispatch, useAppSelector } from "@/app/store";
// import { setCurrentCategoryAutomatically, resetManualSelection } from "@/entities/category/store/categorySlice";

// export const useScrollToCategory = () => {
//   const dispatch = useAppDispatch();
//   const isManuallySelected = useAppSelector(
//     (state) => state.categories.isManuallySelected
//   );
//   const scrollTimeoutRef = useRef<NodeJS.Timeout | null>(null);
//   const isScrollingRef = useRef(false);
//   const lastManualClickTimeRef = useRef<number>(0);
//   const isInitialLoadRef = useRef(true);

//   useEffect(() => {
//     const observerOptions = {
//       root: null,
//       rootMargin: "-20% 0px -60% 0px",
//       threshold: 0
//     };

//     const observerCallback = (entries: IntersectionObserverEntry[]) => {
//       if (isInitialLoadRef.current) {
//         isInitialLoadRef.current = false;
//       } else {
//         const timeSinceClick = Date.now() - lastManualClickTimeRef.current;
//         if (timeSinceClick < 800) {
//           return;
//         }

//         if (isManuallySelected && isScrollingRef.current) {
//           return;
//         }
//       }

//       entries.forEach((entry) => {
//         if (entry.isIntersecting) {
//           const categoryTitle = entry.target.id;
//           if (categoryTitle) {
//             window.history.replaceState(null, "", `#${categoryTitle}`);
          
//             if (!isManuallySelected) {
//               dispatch(setCurrentCategoryAutomatically(categoryTitle));
//               localStorage.setItem("category", categoryTitle);
//             }
//           }
//         }
//       });
//     };

//     const observer = new IntersectionObserver(observerCallback, observerOptions);

//     const handleScroll = () => {
//       isScrollingRef.current = true;
      
//       if (scrollTimeoutRef.current) {
//         clearTimeout(scrollTimeoutRef.current);
//       }

//       scrollTimeoutRef.current = setTimeout(() => {
//         isScrollingRef.current = false;
//         if (isManuallySelected) {
//           dispatch(resetManualSelection());
//         }
//       }, 300);
//     };

//     const handleHashClick = (e: MouseEvent) => {
//       const target = e.target as HTMLElement;
//       const link = target.closest('a[href^="#"]') as HTMLAnchorElement;
//       if (link && link.hash) {
//         lastManualClickTimeRef.current = Date.now();
        
//         isScrollingRef.current = true;
        
//         if (scrollTimeoutRef.current) {
//           clearTimeout(scrollTimeoutRef.current);
//         }

//         scrollTimeoutRef.current = setTimeout(() => {
//           isScrollingRef.current = false;
//           if (isManuallySelected) {
//             dispatch(resetManualSelection());
//           }
//         }, 500);
//       }
//     };

//     const categoryElements = document.querySelectorAll('[id]');
//     categoryElements.forEach((element) => {
//       observer.observe(element);
//     });

//     window.addEventListener("scroll", handleScroll, { passive: true });
//     document.addEventListener("click", handleHashClick, true);

//     return () => {
//       observer.disconnect();
//       window.removeEventListener("scroll", handleScroll);
//       document.removeEventListener("click", handleHashClick, true);
//       if (scrollTimeoutRef.current) {
//         clearTimeout(scrollTimeoutRef.current);
//       }
//     };
//   }, [dispatch, isManuallySelected]);
// };


import { useEffect, useRef } from "react";
import { useAppDispatch, useAppSelector } from "@/app/store";
import { setCurrentCategoryAutomatically, resetManualSelection } from "@/entities/category/store/categorySlice";

export const useScrollToCategory = () => {
  const dispatch = useAppDispatch();
  const isManuallySelected = useAppSelector(
    (state) => state.categories.isManuallySelected
  );
  const scrollTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const isScrollingRef = useRef(false);
  const lastManualClickTimeRef = useRef<number>(0);
  const isInitialLoadRef = useRef(true);
  const hashNavigationTimeoutRef = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
    // Ждем завершения навигации по хешу при загрузке
    const handleInitialHashNavigation = () => {
      if (window.location.hash) {
        // Даем время браузеру выполнить нативную навигацию
        hashNavigationTimeoutRef.current = setTimeout(() => {
          isInitialLoadRef.current = false;
        }, 100);
      } else {
        isInitialLoadRef.current = false;
      }
    };

    // Запускаем после полной загрузки страницы
    if (document.readyState === 'complete') {
      handleInitialHashNavigation();
    } else {
      window.addEventListener('load', handleInitialHashNavigation);
    }

    const observerOptions = {
      root: null,
      rootMargin: "-20% 0px -60% 0px",
      threshold: 0
    };

    const observerCallback = (entries: IntersectionObserverEntry[]) => {
      // На начальной загрузке игнорируем первые срабатывания
      if (isInitialLoadRef.current) {
        return;
      }

      const timeSinceClick = Date.now() - lastManualClickTimeRef.current;
      if (timeSinceClick < 800) {
        return;
      }

      if (isManuallySelected && isScrollingRef.current) {
        return;
      }

      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const categoryTitle = entry.target.id;
          if (categoryTitle) {
            // Обновляем URL только если это не текущий хеш
            if (window.location.hash !== `#${categoryTitle}`) {
              window.history.replaceState(null, "", `#${categoryTitle}`);
            }
            
            if (!isManuallySelected) {
              dispatch(setCurrentCategoryAutomatically(categoryTitle));
              localStorage.setItem("category", categoryTitle);
            }
          }
        }
      });
    };

    const observer = new IntersectionObserver(observerCallback, observerOptions);

    const handleScroll = () => {
      isScrollingRef.current = true;
      
      if (scrollTimeoutRef.current) {
        clearTimeout(scrollTimeoutRef.current);
      }

      scrollTimeoutRef.current = setTimeout(() => {
        isScrollingRef.current = false;
        if (isManuallySelected) {
          dispatch(resetManualSelection());
        }
      }, 300);
    };

    const handleHashClick = (e: MouseEvent) => {
      const target = e.target as HTMLElement;
      const link = target.closest('a[href^="#"]') as HTMLAnchorElement;
      if (link && link.hash) {
        lastManualClickTimeRef.current = Date.now();
        
        isScrollingRef.current = true;
        
        if (scrollTimeoutRef.current) {
          clearTimeout(scrollTimeoutRef.current);
        }

        scrollTimeoutRef.current = setTimeout(() => {
          isScrollingRef.current = false;
          if (isManuallySelected) {
            dispatch(resetManualSelection());
          }
        }, 500);
      }
    };

    // Наблюдаем за элементами с небольшой задержкой
    setTimeout(() => {
      const categoryElements = document.querySelectorAll('[id]');
      categoryElements.forEach((element) => {
        observer.observe(element);
      });
    }, 50);

    window.addEventListener("scroll", handleScroll, { passive: true });
    document.addEventListener("click", handleHashClick, true);

    return () => {
      observer.disconnect();
      window.removeEventListener("scroll", handleScroll);
      document.removeEventListener("click", handleHashClick, true);
      window.removeEventListener('load', handleInitialHashNavigation);
      
      if (scrollTimeoutRef.current) {
        clearTimeout(scrollTimeoutRef.current);
      }
      if (hashNavigationTimeoutRef.current) {
        clearTimeout(hashNavigationTimeoutRef.current);
      }
    };
  }, [dispatch, isManuallySelected]);
};