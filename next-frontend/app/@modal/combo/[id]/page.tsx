import { getComboById } from "@/entities/combo/lib/api";
import { ComboModal } from "@/entities/combo/ui/combo-modal";

interface ComboModalPageProps {
  params: {
    id: string;
  };
}

export default async function ComboModalPage({ params }: ComboModalPageProps) {
  const { id } = params;
  const combo = await getComboById(Number(id));

  if (!combo) {
    return <div>Комбо не найдено</div>;
  }

  return <ComboModal combo={combo} />;
}
