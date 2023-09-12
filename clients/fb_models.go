package clients

import "fmt"

type (
	// WebHookRequest received from Facebook server on webhook, contains messages, delivery reports and/or postbacks.
	WebHookRequest struct {
		Object string                `json:"object,omitempty"`
		Entry  []WebHookRequestEntry `json:"entry,omitempty"`
	}

	// WebHookRequestEntry is an entry in the Facebook web hooks.
	WebHookRequestEntry struct {
		Time      int                          `json:"time,omitempty"`
		ID        string                       `json:"id,omitempty"`
		Messaging []WebHookRequestEntryMessage `json:"messaging,omitempty"`
	}

	// WebHookRequestEntryMessage is a message from user in the Facebook web hook request.
	WebHookRequestEntryMessage struct {
		Timestamp int              `json:"timestamp,omitempty"`
		Message   Message         `json:"message,omitempty"`
		// Delivery  *Delivery        `json:"delivery"`
		// Postback  *Postback        `json:"postback"`
		Recipient User `json:"recipient,omitempty"`
		Sender    User    `json:"sender,omitempty"`
	}

	// MessageRecipient is the recipient data in the Facebook web hook request.
	User struct {
		ID string `json:"id,omitempty"`
	}

	// MessageSender is the sender data in the Facebook web hook request.
	// MessageSender struct {
	// 	ID string `json:"id"`
	// }

	// Message struct for text messaged received from facebook server as part of WebHookRequest struct.
	Message struct {
		Mid        string      `json:"mid,omitempty"`
		// Seq        int         `json:"seq,omitempty"`
		Text       string      `json:"text,omitempty"`
		QuickReply *struct {
			Payload string `json:"payload,omitempty"`
		} `json:"quick_reply,omitempty"`
		Attachments *[]Attachment `json:"attachments,omitempty"`
		Attachment *Attachment `json:"attachment,omitempty"`
	}

	// Attachment is the Facebook messenger message attachment. E.g. buttons.
	Attachment struct {
		Type    string            `json:"type,omitempty"`
		Payload AttachmentPayload `json:"payload,omitempty"`
	}

	// AttachmentPayload is the Facebook messenger message attachment payload.
	AttachmentPayload struct {
		URL string `json:"url,omitempty"`
		// TemplateType string            `json:"template_type"`
		// Text         string            `json:"text"`
		// Buttons      AttachmentButtons `json:"buttons"`
	}

 Response struct {
		Recipient User    `json:"recipient,omitempty"`
		Message   Message `json:"message,omitempty"`
	}

	// AttachmentButtons is the Facebook messenger attachment buttons.
	// AttachmentButtons []AttachmentButton

	// // AttachmentButton is the Facebook messenger attachment button.
	// AttachmentButton struct {
	// 	Type    string `json:"type"`
	// 	Title   string `json:"title"`
	// 	Payload string `json:"payload"`
	// }
	// Delivery struct for delivery reports received from Facebook server as part of WebHookRequest struct.
	// Delivery struct {
	// 	Mids      []string `json:"mids"`
	// 	Seq       int      `json:"seq"`
	// 	Watermark int      `json:"watermark"`
	// }

	// // Postback struct for postbacks received from Facebook server  as part of WebHookRequest struct.
	// Postback struct {
	// 	Payload string `json:"payload"`
	// }

	// SendMessageRequest is a request to send message through FB Messenger
	SendMessageRequest struct {
		MessagingType string           `json:"messaging_type"`
		// Tag           string           `json:"tag,omitempty"`
		RecipientID   User 						`json:"recipient"`
		Message       Message          `json:"message"`
	}

	// APIResponse received from Facebook server after sending the message.
	APIResponse struct {
		MessageID   string    `json:"message_id"`
		RecipientID string    `json:"recipient_id"`
		Error       *APIError `json:"error,omitempty"`
	}

	// APIError received from Facebook server if sending messages failed.
	APIError struct {
		Code      int    `json:"code"`
		FbtraceID string `json:"fbtrace_id"`
		Message   string `json:"message"`
		Type      string `json:"type"`
	}
)

// Error returns Go error object constructed from APIError data.
func (err *APIError) Error() error {
	return fmt.Errorf("FB Error %d: Type %s: %s; FB trace ID: %s", err.Code, err.Type, err.Message, err.FbtraceID)
}