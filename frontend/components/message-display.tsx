"use client";

export default function MessageDisplay({ message }: { message: string }) {
  return (
    <div className="flex flex-col gap-4">
      <div>
        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="message">
          Message
        </label>
        <textarea
          id="message"
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          value={message}
          readOnly={true}
          rows={10}
        />
      </div>
    </div>
  );
}
