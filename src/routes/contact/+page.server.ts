import sgMail from '@sendgrid/mail';

export const actions = {
	default: async ({ request }) => {
		const data = await request.formData();

		const name = data.get('name') ?? 'Anonymous';
		const email = data.get('email') ?? '';
		if (email === '') {
			return {
				status: 400,
				body: 'Email is required!'
			};
		}

		const message = data.get('message') ?? '';
		if (message === '') {
			return {
				status: 400,
				body: 'Message is required!'
			};
		}

		const mail = {
			to: 'contact@asmussen.tech',
			from: email as string,
			subject: `New message from ${name}`,
			text: message as string
		};

		sgMail.setApiKey(process.env.SENDGRID_API_KEY ?? '');
		sgMail
			.send(mail)
			.then(() => {
				console.log(`Sent email from ${email}.`);
			})
			.catch((error) => {
				console.error(error);
			});
	}
};
