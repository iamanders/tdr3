const getTime = async (id) => {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/times/${id}`);

  if (!response.ok) {
    if (response.status === 404) {
      const err = new Error('404Exception');
      throw err;
    } else {
      throw new Error({ name: 'Error', message: response.statusText });
    }
  }

  const time = await response.json();
  return time ?? [];
};

const getTimesBetween = async (startDate, stopDate) => {
  const params = new URLSearchParams({ startDate, stopDate });
  const response = await fetch(`${process.env.VUE_APP_API_URL}/times?${params}`);
  const times = await response.json();

  return times ?? [];
};

const saveTime = async (client, project, code, comment, startDate, stopDate) => {
  const params = {
    client,
    project,
    code,
    comment,
    startedAt: startDate,
    stoppedAt: stopDate,
  };

  const response = await fetch(`${process.env.VUE_APP_API_URL}/times`, {
    method: 'POST',
    body: JSON.stringify(params),
  });

  if (!response.ok) {
    if (response.status === 422) {
      const err = new Error('ValidationErrorException');
      err.validationErrors = await response.json();
      throw err;
    } else {
      throw new Error({ name: 'Error', message: response.statusText });
    }
  }

  return response.json();
};

const updateTime = async (id, client, project, code, comment, startDate, stopDate) => {
  const params = {
    client,
    project,
    code,
    comment,
    startedAt: startDate,
    stoppedAt: stopDate,
  };

  const response = await fetch(`${process.env.VUE_APP_API_URL}/times/${id}`, {
    method: 'PUT',
    body: JSON.stringify(params),
  });

  if (!response.ok) {
    if (response.status === 422) {
      const err = new Error('ValidationErrorException');
      err.validationErrors = await response.json();
      throw err;
    } else {
      throw new Error({ name: 'Error', message: response.statusText });
    }
  }

  return response.json();
};

const deleteTime = async (id) => {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/times/${id}`, {
    method: 'DELETE',
  });

  if (!response.ok) {
    if (response.status === 404) {
      const err = new Error('404Exception');
      throw err;
    } else {
      throw new Error({ name: 'Error', message: response.statusText });
    }
  }

  return JSON.stringify('ok');
};

const stopTimes = async (stopDate) => {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/status/stop`, {
    method: 'PUT',
    body: JSON.stringify({ stoppedAt: stopDate }),
  });

  if (!response.ok) {
    throw new Error({ name: 'Error', message: response.statusText });
  }

  return JSON.stringify('ok');
};

const searchTimes = async (client, project, code, comment, startDate, stopDate) => {
  const params = new URLSearchParams({
    client, project, code, comment, startDate, stopDate,
  });
  const response = await fetch(`${process.env.VUE_APP_API_URL}/times/search?${params}`);
  const times = await response.json();

  return times ?? [];
};

const dataTime = {
  getTime,
  getTimesBetween,
  saveTime,
  updateTime,
  deleteTime,
  stopTimes,
  searchTimes,
};
export { dataTime as default };
