import { HomePage } from "@/pages/home-page";
import { getComboById } from "@/entities/combo/lib/api";
import { ComboModal } from "@/entities/combo/ui/combo-modal";

interface ComboPageProps {
  params: {
    id: string;
  }
}

export default async function ComboPage({ params }: ComboPageProps) {
  const { id } = params;
  const combo = await getComboById(Number(id));

  if (!combo) {
    return <div>Комбо не найдено</div>;
  }

  return (
    <>
      <HomePage />
      <ComboModal combo={combo} />
    </>
  );
}
