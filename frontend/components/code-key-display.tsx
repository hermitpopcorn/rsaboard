"use client";

export default function CodeAndKeyDisplay({
  code,
  privateKey,
}: {
  code: string;
  privateKey: string;
}) {
  return (
    <div className="flex flex-col gap-4">
      <div>
        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="code">
          Message Code
        </label>
        <input
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="code"
          type="text"
          value={code}
          readOnly={true}
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
          id="privateKey"
          value={privateKey}
          readOnly={true}
        />
      </div>
    </div>
  );
}
