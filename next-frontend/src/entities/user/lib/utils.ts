export const normalizePhoneNumber = (phone: string): string => {

  const newPhone = phone.replaceAll(' ', '').replace('+', '').replaceAll('-', '');

  return newPhone;
}