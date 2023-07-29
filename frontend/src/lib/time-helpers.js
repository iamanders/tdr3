import dayjs from 'dayjs';

const formatDuration = (duration) => {
  let formatted = `${Math.floor(dayjs.duration(duration).asHours())}${dayjs.duration(duration).format(':mm')}`;
  if (formatted.length <= 4) {
    formatted = `0${formatted}`;
  }
  return formatted;
};

const durationDiff = (startDate, endDate) => {
  const start = dayjs.utc(startDate);
  const end = endDate !== null ? dayjs.utc(endDate) : dayjs.utc(dayjs().format('YYYY-MM-DD HH:mm:ss'));
  return formatDuration(end.diff(start));
};

const durationDiffMultiple = (times) => {
  const totalDuration = times.reduce((carry, time) => {
    const start = dayjs.utc(time.startedAt);
    const end = time.stoppedAt !== null ? dayjs.utc(time.stoppedAt) : dayjs.utc(dayjs().format('YYYY-MM-DD HH:mm:ss'));

    // eslint-disable-next-line no-param-reassign
    carry += end.diff(start);
    return carry;
  }, 0);
  return formatDuration(totalDuration);
};

const durationDiffGroupedMultiple = (groupedTimes) => {
  let totalDuration = 0;
  groupedTimes.forEach((group) => {
    totalDuration += group.times.reduce((carry, time) => {
      const start = dayjs.utc(time.startedAt);
      const end = time.stoppedAt !== null ? dayjs.utc(time.stoppedAt) : dayjs.utc(dayjs().format('YYYY-MM-DD HH:mm:ss'));

      // eslint-disable-next-line no-param-reassign
      carry += end.diff(start);
      return carry;
    }, 0);
  });

  return formatDuration(totalDuration);
};

const groupTimesByDay = (ungroupedTimes) => {
  const grouped = [];

  ungroupedTimes.forEach((time) => {
    const formattedDate = dayjs(time.startedAt).utc().format('dddd, D MMMM');
    let group = grouped.find((g) => g.formattedDate === formattedDate);
    if (group === undefined) {
      group = {
        formattedDate,
        times: [],
      };
      grouped.push(group);
    }
    group.times.push(time);
  });

  return grouped;
};

const timeHelpers = {
  durationDiff,
  durationDiffMultiple,
  durationDiffGroupedMultiple,
  groupTimesByDay,
};
export { timeHelpers as default };
