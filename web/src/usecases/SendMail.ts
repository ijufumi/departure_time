interface SendMail {
  send(args: { toAddress: string, fromAddress: string, subject: string, body: string }): Promise<any>
}

export default SendMail;