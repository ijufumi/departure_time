interface SendMailRepository {
    send(args: { toAddress: string, subject: string, body: string }): Promise<any>;
}

export default SendMailRepository;