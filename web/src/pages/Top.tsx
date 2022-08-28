import React, { FC } from "react";
import { Stack, Box } from "@mui/material";

interface Props { }

const Top: FC<Props> = () => {
  return (
    <Stack>
      <Box>{"top..."}</Box>
    </Stack>
  );
};

export default Top;
