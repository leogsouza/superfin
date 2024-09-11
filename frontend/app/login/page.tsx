'use client'

import React, { FormEvent } from "react"
import { toast } from 'react-toastify'
import Auth from "@/components/Auth"

const Login = () => {
  const onSubmit = (
    e: FormEvent<HTMLFormElement>,
    formRef: React.RefObject<HTMLFormElement>
  ) => {
    e.preventDefault()
    toast("Login toast", {
      type: "success",
    })
  }

  return <Auth showRemembered onSubmit={onSubmit} />
}