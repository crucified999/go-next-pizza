// // shared/ui/checkout-stepper/simple-checkout-stepper.tsx
// import React from "react";
// import { Check, ShoppingCart, FileText, CheckCircle } from "lucide-react";
// import { cn } from "@/shared/lib/utils";

// export type CheckoutStep = "cart" | "checkout" | "confirmed";

// interface CheckoutStepperProps {
//   currentStep: CheckoutStep;
//   className?: string;
//   onStepClick?: (step: CheckoutStep) => void;
// }

// const steps = [
//   { id: 1, key: "cart" as CheckoutStep, title: "Корзина" },
//   {
//     id: 2,
//     key: "checkout" as CheckoutStep,
//     title: "Оформление заказа",
//   },
//   {
//     id: 3,
//     key: "confirmed" as CheckoutStep,
//     title: "Заказ принят",
//   },
// ];

// export const CheckoutStepper: React.FC<CheckoutStepperProps> = ({
//   currentStep,
//   className,
//   onStepClick,
// }) => {
//   const currentStepIndex = steps.findIndex((step) => step.key === currentStep);

//   return (
//     <div className={cn("w-full max-w-xl mx-auto", className)}>
//       <div className="relative">
//         <div className="relative flex justify-between px-12">
//           {steps.map((step, index) => {
//             const isCompleted = index < currentStepIndex;
//             const isCurrent = index === currentStepIndex;
//             const isPending = index > currentStepIndex;

//             return (
//               <div
//                 key={step.id}
//                 className="flex flex-col items-center relative flex-1"
//               >
//                 <button
//                   onClick={() => onStepClick?.(step.key)}
//                   disabled={!isCompleted && !isCurrent}
//                   className={cn(
//                     "w-8 h-8 rounded-full flex items-center justify-center mb-2 font-bold",
//                     "border-1 border-black bg-white z-20",
//                     isCompleted && "text-black border-1",
//                     isPending && "border-gray-300 text-gray-400"
//                   )}
//                 >
//                   {step.id}
//                 </button>

//                 <h3
//                   className={cn(
//                     "text-sm font-semibold text-black",
//                     isPending && "text-gray-400 opacity-60"
//                   )}
//                 >
//                   {step.title}
//                 </h3>
//               </div>
//             );
//           })}
//         </div>

//         <div className="absolute top-[17px] left-0 right-0 h-0.5 -translate-y-1/2 pointer-events-none">
//           <div
//             className="absolute h-0.5"
//             style={{
//               left: "calc(25% + 15px)",
//               width: "calc(100px)",
//             }}
//           >
//             {currentStepIndex >= 1 && (
//               <div className="absolute inset-0 bg-black" />
//             )}
//           </div>

//           <div
//             className="absolute h-0.5"
//             style={{
//               left: "calc(50% + 30px)",
//               width: "calc(102px)",
//             }}
//           >
//             {currentStepIndex === 1 && (
//               <div
//                 className="absolute inset-0 bg-black opacity-70"
//                 style={
//                   {
//                     "--dash-color": "black",
//                     "--dash-size": "5px",
//                     "--gap-size": "5px",
//                     "--opacity": "0.1",
//                     background: `repeating-linear-gradient(
//                           to right,
//                           var(--dash-color) 0,
//                           var(--dash-color) var(--dash-size),
//                           transparent var(--dash-size),
//                           transparent calc(var(--dash-size) + var(--gap-size))
//                         )`,
//                     opacity: "var(--opacity)",
//                   } as React.CSSProperties
//                 }
//               />
//             )}
//           </div>
//         </div>
//       </div>
//     </div>
//   );
// };

// shared/ui/checkout-stepper/simple-checkout-stepper.tsx
import React from "react";
import { cn } from "@/shared/lib/utils";

export type CheckoutStep = "cart" | "checkout" | "confirmed";

interface CheckoutStepperProps {
  currentStep: CheckoutStep;
  className?: string;
  onStepClick?: (step: CheckoutStep) => void;
}

const steps = [
  { id: 1, key: "cart" as CheckoutStep, title: "Корзина" },
  {
    id: 2,
    key: "checkout" as CheckoutStep,
    title: "Оформление заказа",
  },
  {
    id: 3,
    key: "confirmed" as CheckoutStep,
    title: "Заказ принят",
  },
];

export const CheckoutStepper: React.FC<CheckoutStepperProps> = ({
  currentStep,
  className,
  onStepClick,
}) => {
  const currentStepIndex = steps.findIndex((step) => step.key === currentStep);

  return (
    <div className={cn("w-full max-w-xl mx-auto", className)}>
      <div className="relative">
        <div className="relative flex justify-between px-8 md:px-12">
          {steps.map((step, index) => {
            const isCompleted = index < currentStepIndex;
            const isCurrent = index === currentStepIndex;
            const isPending = index > currentStepIndex;

            return (
              <div
                key={step.id}
                className="flex flex-col items-center relative flex-1"
              >
                <button
                  onClick={() => onStepClick?.(step.key)}
                  disabled={!isCompleted && !isCurrent}
                  className={cn(
                    "w-8 h-8 rounded-full flex items-center justify-center mb-2 font-bold",
                    "border border-black dark:bg-[#101113] dark:border-white dark:text-white bg-white z-20 relative",
                    isCompleted && "text-black",
                    isPending && "border-gray-300 text-gray-400"
                  )}
                >
                  {step.id}
                </button>

                <h3
                  className={cn(
                    "text-sm font-semibold text-black text-center px-1 dark:text-white",
                    isPending && "text-gray-400 opacity-60"
                  )}
                >
                  {step.title}
                </h3>
              </div>
            );
          })}
        </div>

        <div className="absolute top-4 left-0 right-0 h-0.5 -translate-y-1/2 pointer-events-none">
          <div
            className="absolute h-0.5"
            style={{
              left: "calc(25% + 5px)",
              width: "calc(25% - 20px)",
            }}
          >
            {currentStepIndex >= 1 && (
              <div className="absolute inset-0 bg-black dark:bg-white" />
            )}
          </div>

          <div
            className="absolute h-0.5"
            style={{
              left: "calc(50% + 18px)", 
              width: "calc(25% - 20px)",
            }}
          >
            {currentStepIndex === 1 && (
              <div
                className="absolute inset-0"
                style={{
                  background: `repeating-linear-gradient(
                    to right,
                    #000,
                    #000 5px,
                    transparent 5px,
                    transparent 10px
                  )`,
                  opacity: 0.1,
                }}
              />
            )}
          </div>
        </div>
      </div>
    </div>
  );
};