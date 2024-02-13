import classNames from "classnames";

type Option = {
  id: string;
  text: string;
};

const basicClasses =
  "h-full w-full p-4 flex items-center justify-center border-gray-200 focus:ring-4 focus:ring-blue-300 active focus:outline-none";

const activeClasses = "text-gray-900 bg-gray-100";
const inactiveClasses = "bg-white";

export default function MenuTab({
  options,
  activeId,
  activeIdSwitcher,
}: {
  options: Option[];
  activeId: string;
  activeIdSwitcher: (to: string) => void;
}) {
  return (
    <>
      <div className="sm:hidden">
        <label htmlFor="tabs" className="sr-only">
          Select action
        </label>
        <select
          id="tabs"
          className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        >
          {options.map((option) => (
            <option key={option.id} value={option.id}>
              {option.text}
            </option>
          ))}
        </select>
      </div>
      <ul className="hidden text-sm font-medium text-gray-500 rounded-lg shadow sm:flex dark:divide-gray-700 dark:text-gray-400">
        {options.map((option) => {
          const classes = ((active: boolean): string =>
            classNames(basicClasses, active ? activeClasses : inactiveClasses))(
            option.id == activeId,
          );
          return (
            <li className="w-full cursor" key={option.id}>
              <button
                className={classes}
                aria-current="page"
                onClick={() => activeIdSwitcher(option.id)}
              >
                {option.text}
              </button>
            </li>
          );
        })}
      </ul>
    </>
  );
}
