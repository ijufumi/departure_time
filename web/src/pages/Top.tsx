import React, { FC, useState } from "react";
import { Stack, Box, Container, Typography, Grid, Input, Button, Snackbar, Alert } from "@mui/material";
import UsecaseFactory from "../usecases/UsecaseFactory";

interface Props { }

const Top: FC<Props> = () => {
  const [toAddress, setToAddress] = useState<string>("");
  const [fromAddress, setFromAddress] = useState<string>("");
  const [subject, setSubject] = useState<string>("");
  const [body, setBody] = useState<string>("");
  const [successSnackBarOpen, setSuccessSnackBarOpen] = useState<boolean>(false);
  const [errorSnackBarOpen, setErrorSnackBarOpen] = useState<boolean>(false);

  const sendMail = UsecaseFactory.createSendMail();

  const handleSendButton = async () => {
    const result = await sendMail.send({toAddress, fromAddress, subject, body});
    if (result) {
      setToAddress("");
      setFromAddress("");
      setSubject("");
      setBody("");
      setSuccessSnackBarOpen(true);
    } else {
      setErrorSnackBarOpen(true);
    }
  };

  return (
    <Container sx={{ bgcolor: "#f5f5f5", display: "flex", justifyContent: "center", alignItems: "center", minHeight: "100vh", minWidth: "100vw"}}>
      <Box sx={{ bgcolor: "#ffffff", borderRadius: "10px" }} width={"600px"} height={"550px"}>
        <Snackbar open={successSnackBarOpen} onClick={() => setSuccessSnackBarOpen(false)} autoHideDuration={1000}>
          <Alert severity="success" variant="filled">Sending message was successful</Alert>
        </Snackbar>
        <Snackbar open={errorSnackBarOpen} onClick={() => setErrorSnackBarOpen(false)} autoHideDuration={1000}>
          <Alert severity="warning" variant="filled">Sending message was fault by something error</Alert>
        </Snackbar>
        <Stack direction={"column"} sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
          <Box my={2}>
            <Typography variant="h3">{"Send Email"}</Typography>
          </Box>
          <Box sx={{ width: "90%"}}>
            <Grid container spacing={2}>
              <Grid item xs={4}>
                <Typography variant="h5">{"To Address"}</Typography>
              </Grid>
              <Grid item xs={8}>
                <Input type="email" fullWidth={true} value={toAddress} onChange={(e: any) => setToAddress(e.target.value)} />
              </Grid>
              <Grid item xs={4}>
                <Typography variant="h5">{"From Address"}</Typography>
              </Grid>
              <Grid item xs={8}>
                <Input type="email" fullWidth={true} value={fromAddress} onChange={(e: any) => setFromAddress(e.target.value)} />
              </Grid>
              <Grid item xs={4}>
                <Typography variant="h5">{"Subject"}</Typography>
              </Grid>
              <Grid item xs={8}>
                <Input type="text" fullWidth={true} value={subject} onChange={(e: any) => setSubject(e.target.value)} />
              </Grid>
              <Grid item xs={4}>
                <Typography variant="h5">{"Body"}</Typography>
              </Grid>
              <Grid item xs={8}>
                <Input type="text" fullWidth={true} multiline={true} value={body} minRows={10} maxRows={10} onChange={(e: any) => setBody(e.target.value)} />
              </Grid>
              <Grid item xs={12}>
                <Box sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                  <Button variant="outlined" color="success" onClick={handleSendButton}>
                    <Typography variant="button">Send</Typography>
                  </Button>
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Stack>
      </Box>
    </Container>
  );
};

export default Top;
