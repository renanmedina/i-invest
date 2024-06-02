export const parseDMY = (s: string) => {
  let [d, m, y] = s.split(/\D/);
  return [y, m, d].join('-');
};
