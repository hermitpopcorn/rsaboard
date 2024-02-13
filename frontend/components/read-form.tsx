"use client";

import { ChangeEvent, FormEvent, useState } from "react";

export type FormValueTypes = {
  code: string;
  privateKey: string;
};

export default function ReadForm({
  submitHandler,
  formLock,
}: {
  submitHandler: (params) => void;
  formLock: boolean;
}) {
  const [formValues, setFormValues] = useState<FormValueTypes>({
    code: "",
    privateKey: "",
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
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="code">
            Message Code
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            name="code"
            id="code"
            type="text"
            value={formValues.code}
            onChange={handleInput}
          />
        </div>
        <div>
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="privateKey"
          >
            Private Key
          </label>
          <textarea
            className="h-28 shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            name="privateKey"
            id="privateKey"
            value={formValues.privateKey}
            onChange={handleInput}
          />
        </div>
        <div className="flex items-center justify-center">
          <button
            className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Read
          </button>
        </div>
      </div>
    </form>
  );
}
