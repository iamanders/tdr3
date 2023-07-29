const getStatus = async () => {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/status`);

  if (!response.ok) {
    throw new Error({ name: 'Error', message: response.statusText });
  }
  const status = await response.json();
  return status;
};

const dataStatus = { getStatus };
export { dataStatus as default };
