/* eslint-disable react-hooks/exhaustive-deps */
import { useCallback, useEffect, useState } from "react";

const useFetch = (url, options = {}) => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [initialFetch, setInitialFetch] = useState(true); // State to manage first fetch

  const fetchData = useCallback(async () => {
    setLoading(true);
    setError(null);

    try {
      const response = await fetch(url, options);

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }

      const result = await response.json();
      setData(result);
    } catch (error) {
      setError(error);
    } finally {
      setLoading(false);
    }
  }, [url]);

  const refetch = useCallback(async () => {
    return await fetchData();
  }, [fetchData]);

  useEffect(() => {
    if (initialFetch) {
      setInitialFetch(false);
      return;
    }
    fetchData();
  }, [url, fetchData, initialFetch]);

  return { data, loading, error, refetch };
};

export default useFetch;
