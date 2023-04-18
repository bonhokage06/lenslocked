package controllers

import "net/http"

func (f *Faq) Create(r *http.Request) (string, interface{}) {
	questions := []Questions{
		{
			Question: "What is Lenslocked?",
			Answer:   "Lenslocked is a website that allows you to share your photos with the world.",
		},
		{
			Question: "How do I share my photos?",
			Answer:   "You can share your photos by creating an account and uploading your photos.",
		},
		{
			Question: "How do I create an account?",
			Answer:   "You can create an account by clicking the 'Sign Up' link in the top right corner of the page.",
		},
		{
			Question: "How do i contact you?",
			Answer:   `Email us at <a class="text-blue-500 hover:text-blue-800 hover:semibold transition ease-in-out duration-1000" href="mailto:bonhokage06@gmail.com">bonhokage06@gmail.com</a>`,
		},
	}
	return "", FaqResponse{
		Questions: questions,
	}
}
