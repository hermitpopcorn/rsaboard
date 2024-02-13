"use client";

import { ChangeEvent, FormEvent, useState } from "react";

export type FormValueTypes = {
  email: string;
  message: string;
  burn: boolean;
};

export default function PostForm({
  submitHandler,
  formLock,
}: {
  submitHandler: (params) => void;
  formLock: boolean;
}) {
  const [formValues, setFormValues] = useState<FormValueTypes>({
    email: "",
    message: "",
    burn: false,
  });

  const handleInput = (e: ChangeEvent) => {
    const target = e.target as HTMLInputElement;
    setFormValues({
      ...formValues,
      [target.name]: target.value,
    });
  };

  const handleCheck = (e: ChangeEvent) => {
    const target = e.target as HTMLInputElement;
    setFormValues({
      ...formValues,
      [target.name]: !formValues[target.name],
    });
  };

  const handleSubmit = (e: FormEvent) => {
    if (formLock) {
      return;
    }

    e.preventDefault();
    submitHandler(formValues);
  };

  return (
    <form className="bg-white rounded px-8 pt-6 pb-8 mb-4 w-80" onSubmit={handleSubmit}>
      <div className="flex flex-col gap-4">
        <div>
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
            Email
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            name="email"
            id="email"
            type="text"
            value={formValues.email}
            onChange={handleInput}
          />
        </div>
        <div>
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="message">
            Message
          </label>
          <textarea
            className="h-28 shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            name="message"
            id="message"
            value={formValues.message}
            onChange={handleInput}
            maxLength={180}
          />
          <p className="text-sm text-right">{formValues.message.length ?? 0}/180</p>
        </div>
        <div className="flex flex-row gap-2 items-center justify-center">
          <input
            type="checkbox"
            checked={formValues.burn}
            onChange={handleCheck}
            value="1"
            name="burn"
            id="burn"
            className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
          />
          <label className="block text-gray-700 text-sm" htmlFor="burn">
            Burn 5 minutes after first read
          </label>
        </div>
        <div className="flex items-center justify-center">
          <button
            className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Save
          </button>
        </div>
      </div>
    </form>
  );
}
