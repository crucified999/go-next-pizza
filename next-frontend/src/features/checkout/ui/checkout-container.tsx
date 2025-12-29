// features/checkout/ui/checkout-container.tsx
"use client";

import { useState } from "react";
import { CheckoutStepper } from "@/shared/ui/stepper/stepper";
// import { CartSummary } from './cart-summary';
// import { DeliveryForm } from './delivery-form';
// import { OrderConfirmation } from './order-confirmation';
import { ArrowLeft } from "lucide-react";
import { useRouter } from "next/navigation";
import type { CheckoutStep } from "@/shared/ui/stepper/stepper";
import { cn } from "@/shared/lib/utils";

export const CheckoutContainer = () => {
  const [currentStep, setCurrentStep] = useState<CheckoutStep>("checkout");

  const handleStepClick = (step: CheckoutStep) => {
    const steps = ["cart", "checkout", "confirmed"];
    const currentIndex = steps.indexOf(currentStep);
    const targetIndex = steps.indexOf(step);

    if (targetIndex <= currentIndex) {
      setCurrentStep(step);
    }
  };

  return (
    <div className="self-center">
      <div className="max-w-6xl">
        <div>
          <CheckoutStepper
            currentStep={currentStep}
            onStepClick={handleStepClick}
          />
        </div>
      </div>
    </div>
  );
};
