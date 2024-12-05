import { Box } from '@mantine/core';
import dayjs from 'dayjs';
import { Calendar, dayjsLocalizer } from 'react-big-calendar';
import 'react-big-calendar/lib/css/react-big-calendar.css';
const localizer = dayjsLocalizer(dayjs);

const CalendarScreen = () => {
  return (
    <Box>
      <Calendar
        localizer={localizer}
        // events={myEventsList}
        startAccessor="start"
        endAccessor="end"
        style={{ height: 500 }}
      />
    </Box>
  );
};
export default CalendarScreen;
export { CalendarScreen };
