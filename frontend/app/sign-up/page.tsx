'use client'

import React, { FormEvent } from "react"
import { toast } from 'react-toastify'
import Auth from "@/components/Auth"

const Register = () => {
  const onSubmit = (
    e: FormEvent<HTMLFormElement>,
    formRef: React.RefObject<HTMLFormElement>
  ) => {
    e.preventDefault()    
    let formData = {
      email: formRef.current?.email,
      password: formRef.current?.password,
    }

    
    toast("Register toast", {
      type: "success",
    })
  }

  return (
    <Auth 
      onSubmit={onSubmit} 
      title="Sign Up"
      buttonTitle="Register"
      accountInfoText={{
        initialText: "Have an Account?",
        actionLink: "/login",
        actionText: "login",

      }}
    />
  )
}

export default Register;