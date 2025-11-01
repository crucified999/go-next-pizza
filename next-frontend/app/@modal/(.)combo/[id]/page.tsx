import { getComboById } from "@/entities/combo/lib/api";
import { ComboModal } from "@/entities/combo/ui/combo-modal";

export default async function ComboModalPage({
  params,
}: {
  params: { id: string };
}) {
  const { id } = params;
  const combo = await getComboById(Number(id));

  if (!combo) {
    return <div>Комбо не найдено</div>;
  }

  return <ComboModal combo={combo} />;
}
