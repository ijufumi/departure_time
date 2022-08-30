interface SendMail {
  sendMail(args: { toAddress: string, fromAddress: string, subject: string, body: string }): Promise<any>
}

export default SendMail;